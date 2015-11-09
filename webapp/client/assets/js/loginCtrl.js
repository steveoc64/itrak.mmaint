(function(){
  // loginCtrl
  'use strict';

  console.log('loading loginCtrl')

  angular.module('itrak').controller('loginCtrl', function($state, loginState){     

    angular.extend(this, {
      username: '',
      passwd: '',
      login:  function () {
          console.log(this)
          console.log("Calling the login function",this.username,this.passwd)
          loginState.login(this.username, this.passwd)
        }

      })
  });

})();
