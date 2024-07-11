1. Generate APK Conf

curl --location 'https://api.am.wso2.com:9095/api/configurator/1.1.0/apis/generate-configuration' \
--header 'Host: api.am.wso2.com' \
--form 'apiType="GRAPHQL"' \
--form 'definition=@"/Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/patient-info-schema.graphql"'

2. Auth token

curl --location 'https://idp.am.wso2.com:9095/oauth2/token' \
--header 'Host: idp.am.wso2.com' \
--header 'Authorization: Basic NDVmMWM1YzgtYTkyZS0xMWVkLWFmYTEtMDI0MmFjMTIwMDAyOjRmYmQ2MmVjLWE5MmUtMTFlZC1hZmExLTAyNDJhYzEyMDAwMg==' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'scope=apk:api_create'

3. Deploy API using apk-conf file

curl --location 'https://api.am.wso2.com:9095/api/deployer/1.1.0/apis/deploy' \
--header 'Authorization: Bearer AUTH_TOKEN' \
--form 'apkConfiguration=@"/Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/patient-info.apk-conf"' \
--form 'definitionFile=@"/Users/gayangiseneviratne/Documents/APIM_APK/APK/Screencasts/APKScreencasts/GraphQL-APK1.1.0/demo-files/patient-info-schema.graphql"' \
--form 'apiType="GRAPHQL"'

4. Invoke the API

curl --location 'https://default.gw.wso2.com:9095/patientinformation/1.0.0' \
--header 'Host: default.gw.wso2.com' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer AUTH_TOKEN' \
--data '{"query":"query ($id: ID!) {\n    patient (\n        id: $id\n    ) {\n        id\n        name\n        age\n        gender\n        contactInfo {\n            phone\n            email\n            address\n        }\n    }\n}","variables":{"id":"12345"}}'