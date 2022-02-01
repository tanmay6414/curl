node {
    stage('Example') {
        
        checkout([
                                $class                           : 'GitSCM',
                                branches                         : [[name: "*/master"]],
                               extensions: [[
                                    $class: 'SparseCheckoutPaths',
                                    sparseCheckoutPaths: [[
                                        $class: 'SparseCheckoutPath',
                                        path  : 'app/'
                                    ]]
                                ]],
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
    }
    stage('Example1') {
        
        checkout([
                                $class                           : 'GitSCM',
                                branches                         : [[name: "*/master"]],
                               extensions: [[
                                    $class: 'SparseCheckoutPaths',
                                    sparseCheckoutPaths: [[
                                        $class: 'SparseCheckoutPath',
                                        path  : 'chart/curl/'
                                    ]]
                                ]],
                                userRemoteConfigs                : [[
                                                                            url          : 'git@github.com:tanmay6414/curl.git'
                                                                    ]]
                        ])
    }
    stage('ls'){
        sh "ls"
}
