;(function(){
  // sitemgrCtrl
  'use strict';

  console.log('Loading SiteMgr Controller')

  // Controller for the SiteMgr Page
  angular.module('itrak').controller('siteMgrCtrl', function($state, loginState){     
    console.log('Running SiteMgr controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
