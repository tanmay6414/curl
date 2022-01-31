node {
    stage("Approval for Deployment") {
        input {
            message "Should we continue?"
            ok "Yes, we should."
            submitter "ankita"
        }
    }
    stage('Example') {
        sh " chmod +x /Users/tanmay.varade/vibrent/github/jenkins-scripts/jenkins/jobs/a.sh"
        echo "tanmay"
        checkout([
                                $class                           : 'GitSCM',
                                branches                         : [[name: "*/master"]],
                                doGenerateSubmoduleConfigurations: false,
                                extensions                       : [
                                    [
                                        $class             : 'SparseCheckoutPaths',
                                        sparseCheckoutPaths: [[path: "app"]]
                                    ],
                                    [
                                        $class             : 'SparseCheckoutPaths',
                                        sparseCheckoutPaths: [[path: "Jenkinsfile"]]
                                    ] 
                               ],
                                gitTool                          : 'Default',
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
        sh "cat Jenkinsfile"
    }
}
