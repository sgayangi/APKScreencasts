# Introduction

This screencast is about GraphQL API functionalities in APK 1.1.0.

The demo has the following outline.
1. Install WSO2 APK 1.1.0.
2. Deploy backend services.
3. Create and deploy GraphQL APIs in APK.
   1. Using apk-conf file
   2. Using CRs generated from the apk-conf file.

## Installation Prerequisites

1. Kubernetes
2. Helm

## Demo API

The sample API used for this application is for a patient information management system, that handles patient information, their medical appointments and prescriptions.

## Install WSO2 APK 1.1.0 

1. Install WSO2 APK 1.1.0.
```
kubectl create ns apk
helm repo add wso2apk https://github.com/wso2/apk/releases/download/1.1.0
helm repo update
helm install apk wso2apk/apk-helm --version 1.1.0 -n apk
```

2. Verify the deployment
```
kubectl get pods -n apk
```

## Deploy backend services for the Patient Information API

1. Deploy the backend service
```
cd k8s
kubectl apply -f . -n apk
```

2. Verify the deployment
```
kubectl get pods -n apk
```



