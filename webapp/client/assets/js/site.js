;(function(){
  // siteCtrl
  'use strict';

  console.log('Loading Site Controller')

  // Controller for the Site Page
  angular.module('itrak').controller('siteCtrl', function($state, loginState){     
    console.log('Running Site controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
