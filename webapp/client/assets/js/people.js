;(function(){
  // peopleCtrl
  'use strict';

  console.log('Loading People Controller')

  // Controller for the SiteMgr Page
  angular.module('itrak').controller('peopleCtrl', function($state, loginState){     
    console.log('Running People controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
