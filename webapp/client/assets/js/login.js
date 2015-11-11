;(function(){
  // loginCtrl
  'use strict';

  console.log('Loading Login Controller')

  // Remote resource for login / logout transactions
  angular.module('itrak').factory('LoginServer', function($resource, ServerName){
    return $resource(ServerName+'/login',{},{
        login: {method: 'POST'},
        logout: {method: 'DELETE'}
      })
    }
  )

  // List of Roles that users can belong to
  angular.module('itrak').constant('UserRoles',['Worker','Vendor','SiteMgr','Admin'])

  // Service for tracking the Login State
  angular.module('itrak').service('loginState', function($state,$http,UserRoles,LoginServer){

    angular.extend(this, {
      
      // State Variables
      loggedIn: false,
      username: '',     
      role: '',
      token: '',

      login: function(username,passwd) { 
        var vm = this
        LoginServer.login({
          username: username, 
          password: passwd
        },function(retval,r){
          vm.loggedIn = true
          vm.username = retval.Username
          vm.role = retval.Role
          vm.token = retval.Token
          console.log('Success',vm)

          switch (vm.role) {
            case '1':
              $state.go("worker")
              break
            case '3':
              $state.go("vendor")
              break
            case '2':
              $state.go("sitemgr")
              break
            case '100':
              $state.go("admin")
              break
            default:
              $state.go("home")
              break
          }
        },function(){
          console.log('Login Failed')
        })
      },

      logout: function() {
        this.loggedIn = false
        this.username = ''
        this.role = ''
        this.token = ''
        LoginServer.logout()
        console.log("Logged out",this)
      },
    })
  })


  // Controller for the Login Page
  angular.module('itrak').controller('loginCtrl', function($state, loginState){     

    angular.extend(this, {
      username: '',
      passwd: '',
      login:  function () {
          loginState.login(this.username, this.passwd)
        }

      })
  });

})();
