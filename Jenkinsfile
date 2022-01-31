node {
    stage('Example') {
        
        checkout([
            $class           : 'GitSCM',
            branches         : [[name: "master"]],
            userRemoteConfigs: [[
                                   
                                    url          : "git@github.com:tanmay6414/web-technology.git",
                                ]],
            extensions       : [[
                                    $class             : 'SparseCheckoutPaths',
                                    sparseCheckoutPaths: [[
                                              $class: 'SparseCheckoutPath',
                                              path  : "web-technology"
                                    ]]
                                ]],
    ])
        

        sh "ls"
    }
}
