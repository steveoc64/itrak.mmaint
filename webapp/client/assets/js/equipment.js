;(function(){
  // equipmentCtrl
  'use strict';

  console.log('Loading Equipment Controller')

  // Controller for the equipment Page
  angular.module('itrak').controller('equipmentCtrl', function($state, loginState){     
    console.log('Running Equipment controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
    })

  });

})();
