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
                                    ],
                                    [
                                        $class: 'RelativeTargetDirectory', 
                                        relativeTargetDir: "/deploymentValuesFileDir"
                                    ]  
                               ],
                                gitTool                          : 'Default',
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
        sh "cat Jenkinsfile"
        sh "ls /deploymentValuesFileDir"
        sh "cat a.sh"
    }
}
