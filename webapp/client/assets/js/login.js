;(function(){
  // loginCtrl
  'use strict';

  console.log('Loading Login Controller')

  // Remote resource for login / logout transactions
  angular.module('itrak').factory('loginServer', function($resource, ServerName){
    return $resource(ServerName+'/login',{},{
        login: {method: 'POST'},
        logout: {method: 'DELETE'}
      })
    }
  )

  // List of Roles that users can belong to
  angular.module('itrak').constant('UserRoles',['Worker','Vendor','SiteMgr','Admin'])

  // Service for tracking the Login State
  angular.module('itrak').service('loginState', function($state,$location,$http,UserRoles,loginServer){

    angular.extend(this, {
      
      // State Variables
      loggedIn: false,
      username: '',     
      role: '',
      token: '',
      homePage: '',
      menu: [],
      home: function() {
        console.log('Calling home function for ',this.role)
        switch (this.role) {
          case '1':
            this.homePage = 'worker'
            this.menu = [
              {sref: "worker.dashboard", title: "Dashboard"},
              {sref: "worker.estop", title: "BreakDown !"},
              {sref: "worker.pstop", title: "Preventative"},
              {sref: "worker.equip", title: "Equipment"},
              {sref: "worker.workorder", title: "WorkOrders"},
              {sref: "worker.task", title: "Tasks"},
              {sref: "worker.spares", title: "Spares"},
              {sref: "worker.reports", title: "Reports"}
            ]
            break
          case '3':
            this.homePage = 'vendor'
            break
          case '2':
            this.homePage = 'sitemgr'
            break
          case '100':
            this.homePage = 'admin'
            this.menu = [
              {sref: "admin.dashboard", title: "Dashboard"},
              {sref: "admin.people", title: "People"},
              {sref: "admin.site", title: "Sites"},
              {sref: "admin.equipment", title: "Equipment"},
              {sref: "admin.workorder", title: "WorkOrders"},
              {sref: "admin.task", title: "Tasks"}
            ]
            break
          default:
            this.homePage = 'login'
            break
        }
        $state.go(this.homePage)
      },
      login: function(username,passwd) { 
        var vm = this
        loginServer.login({
          username: username, 
          password: passwd
        },function(retval,r){
          vm.loggedIn = true
          vm.username = retval.Username
          vm.role = retval.Role
          vm.token = retval.Token
          console.log('Success',vm)
          vm.home()
        },function(){
          console.log('Login Failed')
        })
      },

      logout: function() {
        this.loggedIn = false
        this.username = ''
        this.role = ''
        this.token = ''
        this.homePage = 'login'
        loginServer.logout()
        console.log("Logged out",this)
        $location.path('/')
      },
    })
  })


  // Controller for the Login Page
  angular.module('itrak').controller('loginCtrl', function($state, loginState){     

    angular.extend(this, {
      username: '',
      passwd: '',
      loginState: loginState,
      login:  function () {
          loginState.login(this.username, this.passwd)
        }

      })
  });

})();
