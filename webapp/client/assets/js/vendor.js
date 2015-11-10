;(function(){
  // vendorCtrl
  'use strict';

  console.log('Loading Vendor Controller')

  // Controller for the Vendor Page
  angular.module('itrak').controller('vendorCtrl', function($state, loginState){     
    console.log('Running Vendor controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
