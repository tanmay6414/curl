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
}
