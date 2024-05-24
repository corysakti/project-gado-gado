node('RITA_APP_SIT') {
    stage('Git Repo') {
        //git credentialsId: 'jenkinsgit1', url: 'jenkinsgit@10.0.144.197:/gitdata/BIMA/d3-api.git'
        //git branch: 'live-production', credentialsId: 'jenkinsgit1', url: 'jenkinsgit@10.0.144.197:/gitdata/BIMA/d3-api.git'
        //git branch: params.GITBRANCH, credentialsId: 'hamditmforgitlab', url: 'https://gitlab.com/hamdimobile/iot.git'
        git branch: params.GITBRANCH, credentialsId: 'jenkinsforgitlabce', url: 'http://10.64.17.56/vas1/iot.git'

    }
    stage('Build and Test') {
        sh "mvn clean package -DskipTests"
        echo "${currentBuild.projectName}"
        echo "BUILD_TAG = ${BUILD_TAG}"
        echo "JOB_NAME = ${JOB_NAME}"
        echo "JOB_BASE_NAME = ${JOB_BASE_NAME}"
    }

    stage('Deploy') {
        sh "pwd"
        PROJECT_PATH = sh(script: 'pwd', , returnStdout: true).trim()
        TARGET_WAR = sh(script: 'ls target/iot*.war', , returnStdout: true).trim()
        sh "TARGET_WAR=${TARGET_WAR}"
        sh "PROJECT_PATH=${PROJECT_PATH}"
        sh "ssh iotuser@10.64.25.122 \" mkdir -p /tmp/target/ \""
        sh "scp ${PROJECT_PATH}/${TARGET_WAR} iotuser@10.64.25.122:/tmp/target/"
        sh "ssh iotuser@10.64.25.122 \" /home/iotuser/jenkins.sh /tmp/${TARGET_WAR} \""
    }

}
