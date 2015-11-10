;(function(){
  // adminCtrl
  'use strict';

  console.log('Loading Admin Controller')

  // Controller for the Admin Page
  angular.module('itrak').controller('adminCtrl', function($state, loginState){     
    console.log('Running Admin controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
