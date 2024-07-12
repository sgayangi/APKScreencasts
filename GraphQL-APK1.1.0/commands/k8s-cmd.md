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