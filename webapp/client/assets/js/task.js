;(function(){
  // taskCtrl
  'use strict';

  console.log('Loading Task Controller')

  // Controller for the SiteMgr Page
  angular.module('itrak').controller('taskCtrl', function($state, loginState){     
    console.log('Running Task controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
