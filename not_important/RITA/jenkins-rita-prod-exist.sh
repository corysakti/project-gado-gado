node('RITA_SIT_13') {
    stage('Preparation') {
        if (JBOSS_NODE=='') {
            sh "echo ERROR : No JBOSS_NODE selected !!!"
            sh "exit 1"
        }

        currentBuild.displayName = BUILD_ID+":"+GITBRANCH
        currentBuild.description = "Building from Git branch: "+GITBRANCH+" on node : " + JBOSS_NODE

		echo "Deployment for ${GITBRANCH}"
		echo "${JBOSS_NODE.tokenize(',').toString()}"

    }
    stage('Git Repo') {
        //git branch: params.GITBRANCH, credentialsId: 'hamditmforgitlab', url: 'https://gitlab.com/hamdimobile/rita2-api.git'
        git branch: params.GITBRANCH, credentialsId: 'jenkinsforgitlabce', url: 'http://10.64.17.56/vas1/rita2-api.git'

    }
    stage('Build and Test') {
        sh "cd rita2-service"
        sh "JAVA_HOME=/app/jdk8 && mvn -version"
        sh "JAVA_HOME=/app/jdk8 && mvn clean package -DskipTests"
        sh "cd .."
        echo "${currentBuild.projectName}"
        echo "BUILD_TAG = ${BUILD_TAG}"
        echo "JOB_NAME = ${JOB_NAME}"
        echo "JOB_BASE_NAME = ${JOB_BASE_NAME}"
    }

    stage('Deploy') {
        PROJECT_PATH = sh(script: 'pwd', , returnStdout: true).trim()
        TARGET_WAR = sh(script: 'ls rita2-service/target/rita2*.war', , returnStdout: true).trim()
        WAR_FILE = sh(script: 'basename rita2-service/target/rita2*.war', , returnStdout: true).trim()

        TARGET_HOST = JBOSS_NODE
        TARGET_PORT = "9990"
        JBOSS_USER = "jenkins"
        JBOSS_PASSWORD = "jenkins123"

        sh "PROJECT_PATH=${PROJECT_PATH}"
        sh "TARGET_WAR=${TARGET_WAR}"
        sh "WAR_FILE=${WAR_FILE}"

        //sh "TARGET_HOST=${TARGET_HOST}"
        sh "TARGET_PORT=${TARGET_PORT}"
        sh "JBOSS_USER=${JBOSS_USER}"
        sh "JBOSS_PASSWORD=${JBOSS_PASSWORD}"

		def APP_NODE = JBOSS_NODE.tokenize(',');
		for (i in 0..<APP_NODE.size()) { APP_NODE[i] = APP_NODE[i].trim() }
		APP_NODE = APP_NODE.unique(false);
		echo "List of APP_NODE : "
		echo APP_NODE.toString()

        for(int i = 0; i < APP_NODE.size(); i++) {
			def s = APP_NODE.get(i).trim()
			def TARGET_HOST = "${s}"
			echo "Deployment started on ${TARGET_HOST}"
			try {
				if(TARGET_HOST != null) {
					//sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c controller=${TARGET_HOST}:${TARGET_PORT} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" /deployment=*/:read-resource(recursive=false,proxies=true,include-runtime=true,include-defaults=true)\" "

					echo "Previous Enabled WAR files "
					sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c controller=${TARGET_HOST}:${TARGET_PORT} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" deployment-info --name=*war \"  "

					echo "Undeploy Previous WAR file"
					sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c controller=${TARGET_HOST}:${TARGET_PORT} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" undeploy rita2-*.war --keep-content \" "

					echo "Deploy New WAR file"
					sh "sh /app/EAP-7.0.0/bin/jboss-cli.sh -c controller=${TARGET_HOST}:${TARGET_PORT} --user=${JBOSS_USER} --password=${JBOSS_PASSWORD} --command=\" deploy ${PROJECT_PATH}/${TARGET_WAR} \" "

					echo "Deployment completed on  ${TARGET_HOST}"

				}
			} catch (Exception e){
				echo e.getMessage()
				echo "Failed to Deploy application - "+s
				ErrorFlag='Y'
				FailedApps=FailedApps+"\n"+s
			}
		}

    }

}
