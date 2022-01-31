node {
    stage('Example') {
     
        
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
        sh "ls"
        sh "cat a.sh"
    }
}
