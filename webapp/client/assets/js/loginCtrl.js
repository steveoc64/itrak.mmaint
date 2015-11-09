(function(){
  // loginCtrl
  'use strict';

  console.log('loading loginCtrl')

  angular.module('itrak').controller('loginCtrl', function($scope, $state, loginState){     

    angular.extend($scope, {
      login: function() {
        console.log("Calling the login function",$scope.username,$scope.password)
        loginState.login()
        $state.go('home')
      }
    })  
  });

})();
