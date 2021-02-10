package logr_test

import (
	"context"
	"testing"

	"github.com/gmaliar/logr"
	"github.com/gmaliar/logr/logrfakes"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogr(t *testing.T) {
	l := logr.WithCtx(context.Background())
	assert.NotNil(t, l)
}

func TestLogrWithDifferentLogger(t *testing.T) {
	l := logr.NewWithLogger(logrus.StandardLogger()).WithContext(context.Background())
	assert.NotNil(t, l)
}

func TestLogrWithZipkinContext(t *testing.T) {
	tracer := mocktracer.New()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("test_logger")
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	fl := &logrfakes.FakeFieldLogger{}
	fl.WithFieldsReturns(logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}))

	l := logr.WithCtx(ctx)
	l.Logger = fl
	l.Infof(":aasdasd")
	assert.NotNil(t, l)
}

func TestLogrWithDifferentLoggerZipkinContext(t *testing.T) {
	tracer := mocktracer.New()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("test_logger")
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	fl := &logrfakes.FakeFieldLogger{}
	fl.WithFieldsReturns(logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}))

	l := logr.NewWithLogger(logrus.StandardLogger()).WithContext(ctx)
	l.Logger = fl
	l.Infof(":aasdasd")
	assert.NotNil(t, l)
}
