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

## Deploying the API

### Method 1 - Using the APK-Conf file
1. Generate the APK-Conf file using the GraphQL SDL file.
```
curl --location 'https://127.0.0.1:9095/api/configurator/1.1.0/apis/generate-configuration' \
--header 'Host: api.am.wso2.com' \
--form 'apiType="GRAPHQL"' \
--form 'definition=@"{location-to-sdl-file}"'
```

2. Copy the generated value and paste it in a .apk-conf file. Fill in the fields for name, basepath, version and production or sandbox endpoint configuration.

3. Generate the authentication token. You will need this value to continue with the demo.

```
curl --location 'https://idp.am.wso2.com:9095/oauth2/token' \
--header 'Host: idp.am.wso2.com' \
--header 'Authorization: Basic NDVmMWM1YzgtYTkyZS0xMWVkLWFmYTEtMDI0MmFjMTIwMDAyOjRmYmQ2MmVjLWE5MmUtMTFlZC1hZmExLTAyNDJhYzEyMDAwMg==' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'scope=apk:api_create'
```

4. Deploy the API. Replace ACCESS_TOKEN with the token generated in the previous step.

curl --location 'https://api.am.wso2.com:9095/api/deployer/1.1.0/apis/deploy' \
--header 'Authorization: Bearer ACCESS_TOKEN' \
--form 'apkConfiguration=@"{location-to-apk-conf-file}"' \
--form 'definitionFile=@"{location-to-graphql-file}"' \
--form 'apiType="GRAPHQL"'

