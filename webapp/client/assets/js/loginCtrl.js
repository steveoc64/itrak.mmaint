(function(){
  // loginCtrl
  'use strict';

  console.log('loading loginCtrl')

  angular.module('itrak').controller('loginCtrl', function($state, loginState){     

    this.login = login
    
    function login() {
        console.log(this)
        console.log("Calling the login function",this.username,this.password)
        loginState.login(this.username, this.passwd)
        $state.go('home')
    }

    angular.extend(this, {
      username: '',
      passwd: ''
      })
  });

})();
