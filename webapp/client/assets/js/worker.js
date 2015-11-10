;(function(){
  // workerCtrl
  'use strict';

  console.log('Loading Worker Controller')

  // Remote resource for login / logout transactions
  angular.module('itrak').factory('WorkerServer', function($resource, ServerName){
    return $resource(ServerName+'/worker')
  })


  // Controller for the Worker Page
  angular.module('itrak').controller('workerCtrl', function($state, loginState){     
    console.log('Running Worker controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
