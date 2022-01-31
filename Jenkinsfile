node {
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
                                        sparseCheckoutPaths: [[path: "${cluster}/workloads/keycloak/${stack}.values.yaml"]]
                                    ],
                                    [
                                        $class: 'RelativeTargetDirectory', 
                                        relativeTargetDir: deploymentValuesFileDir
                                    ]                              
                               ],
                                gitTool                          : 'Default',
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
        sh "cat app/main.go"
    }
}
