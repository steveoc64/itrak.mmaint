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
      site: 0,
      siteName: '',
      homePage: '',
      menu: [],
      home: function() {
        console.log('Calling home function for Role =',this.role)
        switch (this.role) {
          case '1':
          case 1:
            this.homePage = 'worker'
            this.menu = [
              {sref: "worker.dashboard", title: "Dashboard"},
              {sref: "worker.estop", title: "BreakDown !"},
              {sref: "worker.pstop", title: "Preventative"},
              {sref: "worker.equip", title: "Equipment"},
              {sref: "worker.task", title: "Tasks"},
              {sref: "worker.spares", title: "Spares"},
              {sref: "worker.reports", title: "Reports"}
            ]
            console.log('we are here with',this)
            break
          case '3':
          case 3:
            this.homePage = 'vendor'
            this.menu = [
              {sref: "vendor.dashboard", title: "Dashboard"},
              {sref: "vendor.equip", title: "Equipment"},
              {sref: "vendor.workorder", title: "WorkOrders"},
              {sref: "vendor.spares", title: "Spares"},
              {sref: "vendor.reports", title: "Reports"}
            ]
            break
          case '2':
          case 2:
            this.homePage = 'sitemgr'
            this.menu = [
              {sref: "sitemgr.dashboard", title: "Dashboard"},
              {sref: "sitemgr.estop", title: "BreakDown !"},
              {sref: "sitemgr.pstop", title: "Preventative"},
              {sref: "sitemgr.equip", title: "Equipment"},
              {sref: "sitemgr.task", title: "Tasks"},
              {sref: "sitemgr.workorder", title: "WorkOrders"},
              {sref: "sitemgr.spares", title: "Spares"},
              {sref: "sitemgr.reports", title: "Reports"}
            ]
            break
          case '100':
          case 100:
            this.homePage = 'admin'
            this.menu = [
              {sref: "admin.people", title: "People"},
              {sref: "admin.site", title: "Sites"},
              {sref: "admin.equipment", title: "Equipment"},
              {sref: "admin.spares", title: "Spares"},
              {sref: "admin.consumables", title: "Consumables"},
              {sref: "admin.equiptypes", title: "Equipment Types"},
              {sref: "admin.task", title: "Tasks"},
              {sref: "admin.workorder", title: "WorkOrders"},
              {sref: "admin.vendor", title: "Vendors"},
            ]
            break
          default:
            this.homePage = 'login'
            break
        }
        console.log('advancing to',this.homePage)
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
          vm.site = retval.Site
          vm.siteName = retval.SiteName
          console.log('Success',vm,retval)
          console.log('Set sitename to ',vm.siteName)
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
