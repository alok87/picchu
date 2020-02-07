package plan

import (
	"context"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/go-logr/logr"
	picchuv1alpha1 "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1"
	"go.medium.engineering/picchu/pkg/plan"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeleteSLORules struct {
	App       string
	Namespace string
}

func (p *DeleteSLORules) Apply(ctx context.Context, cli client.Client, scalingFactor float64, log logr.Logger) error {
	prlist := &monitoringv1.PrometheusRuleList{}

	opts := &client.ListOptions{
		Namespace: p.Namespace,
		LabelSelector: labels.SelectorFromSet(map[string]string{
			picchuv1alpha1.LabelApp: p.App,
		}),
	}

	if err := cli.List(ctx, prlist, opts); err != nil {
		log.Error(err, "Failed to delete SLO Rules")
		return err
	}

	if prlist.Items != nil {
		for _, sm := range prlist.Items {
			err := cli.Delete(ctx, sm)
			if err != nil && !errors.IsNotFound(err) {
				plan.LogSync(log, "deleted", err, sm)
				return err
			}
			if err == nil {
				plan.LogSync(log, "deleted", err, sm)
			}
		}
	}
	return nil
}