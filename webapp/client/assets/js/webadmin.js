;(function(){
  // adminCtrl
  'use strict';

  console.log('Loading Admin Controller')

  // Controller for the Admin Page
  angular.module('itrak').controller('adminCtrl', function(
    $state, 
    loginState,    
    People,Sites,Equipment,WorkOrders,Tasks,Vendors,Spares,Roles    
    ){     
    console.log('Running Admin controller')

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
      menuSelection: '',

      // Beware - Black Magik in here to call auto resolver function
      // - looks for a function with _XXXXX, where XXXXX = name of the
      // sub UI-router view
      go: function(menu) {
        var cstr = menu.sref.split(".")
        if (cstr.length > 1) {
          var fn = this['_'+cstr[1]]
          if (fn) {
            fn.call(this)
          }
        }
        this.menuSelection = menu.title
        $state.go(menu.sref)
      },

      getSiteURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },

      // Called when initialising various screens, see
      // the nasty code above in the 'go' function
      _people: function() {
        this.getPeople()
        this.getSites()
        this.getRoles()
      },
      _site: function() {
        this.getSites()
      },
      _equipment: function() {
        this.getEquipment()
        this.getSites()
      },

      getPeople: function() {
        console.log('Loading People List')
        this.people = People.query()
      },
      getSites: function() {
        console.log('Loading Sites List')
        this.sites = Sites.query()
      },
      getRoles: function() {
        console.log('loading Roles List')
        this.roles = Roles.query()
      },
      getEquipment: function() {
        console.log('Loading Equipment List')
        this.equipment = Equipment.query()
      },


    });
    
  });

})();
