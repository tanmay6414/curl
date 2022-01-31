node {
    stage('Example') {
        
        checkout([
            $class           : 'GitSCM',
            branches         : [[name: "master"]],
            userRemoteConfigs: [[
                                   
                                    url          : "git@github.com:tanmay6414/computational-intelligence.git",
                                ]],
            extensions       : [[
                                    $class             : 'SparseCheckoutPaths',
                                    sparseCheckoutPaths: [[
                                              $class: 'SparseCheckoutPath',
                                              path  : "computational-intelligence"
                                    ]]
                                ]],
    ])
        

        sh "ls web-technology"
    }
}
