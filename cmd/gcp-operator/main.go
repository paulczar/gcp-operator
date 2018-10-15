package main

import (
	"context"
	"runtime"
	"time"

	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	stub "github.com/paulczar/gcp-operator/pkg/stub"

	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()

	sdk.ExposeMetricsPort()

	resource := "cloud.google.com/v1alpha1"
	kinds := []string{"Instance", "Address"}
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("failed to get watch namespace: %v", err)
	}
	resyncPeriod := time.Duration(5) * time.Second
	for _, kind := range kinds {
		logrus.Infof("Watching %s, %s, %s, %d", resource, kind, namespace, resyncPeriod)
		sdk.Watch(resource, kind, namespace, resyncPeriod)
	}
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
