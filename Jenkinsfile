node {
    stage('Example') {
        input {
            message "Should we continue?"
            ok "Yes, we should."
            submitter "ankita"
        }
        
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
