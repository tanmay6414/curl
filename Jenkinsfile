node {
    stage('Example') {
        
        checkout([
                                $class                           : 'GitSCM',
                                branches                         : [[name: "*/master"]],
                                extensions                       : [
                                    [
                                        $class             : 'SparseCheckoutPaths',
                                        sparseCheckoutPaths: [[path: "app"]]
                                    ],
                                    [
                                        $class             : 'SparseCheckoutPaths',
                                        sparseCheckoutPaths: [[path: "a.sh"]]
                                    ],
               
                               ],
                                gitTool                          : 'Default',
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
//         sh "cat Jenkinsfile"
//         sh "ls deploymentValuesFileDir"
//         sh "cat a.sh"
        sh "ls"
    }
}
