1. Generate the CRs using the APK-Conf file

curl --location 'https://api.am.wso2.com:9095/api/configurator/1.1.0/apis/generate-k8s-resources' \
--header 'Host: api.am.wso2.com' \
--form 'apkConfiguration=@"/Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/patient-info.apk-conf"' \
--form 'definitionFile=@"/Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/patient-info-schema.graphql"' \
--form 'apiType="REST"' -k --output /Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/Patient.zip

2. Unzip the zip file

unzip /Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/Patient.zip -d /Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/PatientCRDs

3. Apply the CRDs

cd /Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/PatientCRDs
cd {directory}
kubectl apply -f . -n apk

4. Auth token

curl --location 'https://idp.am.wso2.com:9095/oauth2/token' \
--header 'Host: idp.am.wso2.com' \
--header 'Authorization: Basic NDVmMWM1YzgtYTkyZS0xMWVkLWFmYTEtMDI0MmFjMTIwMDAyOjRmYmQ2MmVjLWE5MmUtMTFlZC1hZmExLTAyNDJhYzEyMDAwMg==' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'scope=apk:api_create'

5. Invoke the API

curl --location 'https://default.gw.wso2.com:9095/patientinformation/1.0.0' \
--header 'Host: default.gw.wso2.com' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsICJ0eXAiOiJKV1QiLCAia2lkIjoiZ2F0ZXdheV9jZXJ0aWZpY2F0ZV9hbGlhcyJ9.eyJpc3MiOiJodHRwczovL2lkcC5hbS53c28yLmNvbS90b2tlbiIsICJzdWIiOiI0NWYxYzVjOC1hOTJlLTExZWQtYWZhMS0wMjQyYWMxMjAwMDIiLCAiYXVkIjoiYXVkMSIsICJleHAiOjE3MjA3NjA3NTQsICJuYmYiOjE3MjA3NTcxNTQsICJpYXQiOjE3MjA3NTcxNTQsICJqdGkiOiIwMWVmNDAwNC0wODk3LTFmZDYtODZlNi1kMWU4YWU2MjZkYmIiLCAiY2xpZW50SWQiOiI0NWYxYzVjOC1hOTJlLTExZWQtYWZhMS0wMjQyYWMxMjAwMDIiLCAic2NvcGUiOiJhcGs6YXBpX2NyZWF0ZSJ9.ILKlO3RnYNON34NAMg7vpq5WiZNlo7SrgiiQVNgwB-K7HE-840pSjnwJWLeQ5DtX6HBGPdk7fTHwp-HOp87G_fo7HhQdFjWKeFF4CQlErypXUl-JM5bOWXPeCsdVGhfVUVlaoeYtIvpB68QT6-377VwAaUL2FJxkq2YouZDTAIRKjLHKOu--p_3bTmSrnpPROSxQTQl-3TWwH_YrejQeqJ44_RIV3rnyRoMPHKpnaGG4wxTOo0yuW2gsKKDlTh_F87U5rrZKzbgR_rrZNfcKnJSd1AWP2g6i3r1rtqh_x_4Dq16gLj7IEehMD3LEFCIKjzNnZ_20tdI0GzDWqWtR5A' \
--data '{"query":"query ($id: ID!) {\n    patient (\n        id: $id\n    ) {\n        id\n        name\n        age\n        gender\n        contactInfo {\n            phone\n            email\n            address\n        }\n    }\n}","variables":{"id":"12345"}}'