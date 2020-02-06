package plan

import (
	"context"
	"fmt"

	slov1alpha1 "github.com/Medium/service-level-operator/pkg/apis/monitoring/v1alpha1"
	"github.com/go-logr/logr"
	picchuv1alpha1 "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1"
	"go.medium.engineering/picchu/pkg/plan"
	prometheus "go.medium.engineering/picchu/pkg/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SyncServiceLevels struct {
	App                         string
	Namespace                   string
	Labels                      map[string]string
	ServiceLevelObjectiveLabels picchuv1alpha1.ServiceLevelObjectiveLabels
	ServiceLevelObjectives      []*picchuv1alpha1.ServiceLevelObjective
}

func (p *SyncServiceLevels) Apply(ctx context.Context, cli client.Client, scalingFactor float64, log logr.Logger) error {
	serviceLevels, err := p.serviceLevels()
	if err != nil {
		return err
	}
	if len(serviceLevels) > 0 {
		for _, sl := range serviceLevels {
			if err := plan.CreateOrUpdate(ctx, log, cli, sl); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *SyncServiceLevels) serviceLevels() ([]*slov1alpha1.ServiceLevel, error) {
	sl := []*slov1alpha1.ServiceLevel{}
	slos := []slov1alpha1.SLO{}

	for _, slo := range p.ServiceLevelObjectives {
		if slo.Enabled {
			config := SLOConfig{
				SLO:    slo,
				App:    p.App,
				Name:   sanitizeName(slo.Name),
				Labels: p.ServiceLevelObjectiveLabels,
			}
			serviceLevelObjective := config.serviceLevelObjective()
			serviceLevelObjective.ServiceLevelIndicator.SLISource.Prometheus = config.sliSource()
			slos = append(slos, *serviceLevelObjective)
		}
	}

	serviceLevel := &slov1alpha1.ServiceLevel{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.App,
			Namespace: p.Namespace,
			Labels:    p.Labels,
		},
		Spec: slov1alpha1.ServiceLevelSpec{
			ServiceLevelName:       p.App,
			ServiceLevelObjectives: slos,
		},
	}
	sl = append(sl, serviceLevel)

	return sl, nil
}

func (p *SLOConfig) sliSource() *slov1alpha1.PrometheusSLISource {
	source := &slov1alpha1.PrometheusSLISource{
		ErrorQuery: p.serviceLevelErrorQuery(),
		TotalQuery: p.serviceLevelTotalQuery(),
	}
	return source
}

func (s *SLOConfig) serviceLevelObjective() *slov1alpha1.SLO {
	labels := make(map[string]string)
	for k, v := range s.Labels.ServiceLevelLabels {
		labels[k] = v
	}

	for k, v := range s.SLO.ServiceLevelObjectiveLabels.ServiceLevelLabels {
		labels[k] = v
	}

	if s.Tag != "" {
		labels[prometheus.TagLabel] = s.Tag
	}

	slo := &slov1alpha1.SLO{
		Name:                         s.Name,
		AvailabilityObjectivePercent: s.SLO.ObjectivePercent,
		Description:                  s.SLO.Description,
		Disable:                      false,
		Output: slov1alpha1.Output{
			Prometheus: &slov1alpha1.PrometheusOutputSource{
				Labels: labels,
			},
		},
	}
	return slo
}

func (s *SLOConfig) serviceLevelTotalQuery() string {
	return fmt.Sprintf("sum(%s)", s.totalQuery())
}

func (s *SLOConfig) serviceLevelErrorQuery() string {
	return fmt.Sprintf("sum(%s)", s.errorQuery())
}
