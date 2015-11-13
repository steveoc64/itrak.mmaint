;(function(){
  // adminCtrl
  'use strict';

  console.log('Loading Admin Controller')

  // Controller for the Admin Page
  angular.module('itrak').controller('adminCtrl', function(
    $state, $stateParams,
    loginState,    
    People,Sites,Equipment,WorkOrders,Tasks,Vendors,Spares,Roles    
    ){     
    console.log('Running Admin controller',$stateParams)

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
      menuSelection: '',
      params: $stateParams,
      selectedSites: [],
      filteredEquips: [],

      // Beware - Black Magik in here to call auto resolver function
      // - looks for a function with _XXXXX, where XXXXX = name of the
      // sub UI-router view
      goMenu: function(menu) {
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

      siteSelected: function(id) {
        console.log('Toggled site',id)
        this.rebuildFilteredEquip()
      },
      rebuildFilteredEquip: function() {
        console.log('Rebuilding filtered equip list from',this.selectedSites)
        this.filteredEquips = []
        angular.forEach(this.equipment, function(v,k){
          if (this.selectedSites[v.location]) {
            this.filteredEquips.push(v)
          }
        },this)
      },
      getSiteURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },

      // Called when initialising various screens, see
      // the nasty code above in the 'goMenu' function
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

      viewEquipmentDetails: function(id) {
        console.log("Getting details for id",id)
        this.equipmentDetails = this.getEquipmentDetails(id)
        if (this.equipmentDetails) {
          console.log('found',this.equipmentDetails)
          $state.go("admin.equipment.details",{"id": id})
        } else {
          alert("Invalid Equipment ID "+id)
        }
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
      getEquipmentDetails: function(id) {
        var e = this.equipment
        console.log('Looking for',id)
        for (var i = 0; i < e.length; i++) {
          if (id === e[i].id) {
            console.log('Found',e[i])
            return e[i]
          }
        };
        return null
      },

    });
    
  });

})();
