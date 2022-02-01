node {
    stage('Example') {
        
        checkout([
                                $class                           : 'GitSCM',
                                branches                         : [[name: "*/master"]],
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
    
    
        
        
        sh 'chmod +x a.sg && ./a.sh'
    }
}
