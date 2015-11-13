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
      siteEquips: [],
      mainSiteEquips: [],
      equipmentTypes: [
        {id: 1, name: 'Consumables'},
        {id: 2, name: 'Mech'},
        {id: 3, name: 'Spare Parts'},
      ],

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
      toggleAllSites: function() {
        if (!this.allSites) {
          // ALL is NOT selected, so turn them all off
          this.selectedSites = []
        } else {
          // Turn ALL on
          this.selectedSites = []
          if (this.siteSelected.length) {
            angular.forEach(this.sites, function(v,k){
              this.selectedSites[v.id] = true
            },this)
          } 
        }
        this.rebuildFilteredEquip()
      },
      rebuildFilteredEquip: function() {
        //console.log('Rebuilding filtered equip list from',this.selectedSites)
        this.siteEquips = []

        // First, get a list of equipments for this site
        angular.forEach(this.equipment, function(v,k){
          if (this.selectedSites[v.location]) {
            this.siteEquips.push(v)
          }
        },this)

        // Now, create a sub-list which is just the parent items on this site
        this.mainSiteEquips = []
        angular.forEach(this.siteEquips, function(v,k){
          if (v.parent_id == 0) {
            this.mainSiteEquips.push(v)
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
        this.getVendors()
      },

      viewEquipmentDetails: function(id) {
        console.log("Getting details for id",id)
        if (this.equipmentDetails(id)) {
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
      getVendors: function() {
        console.log('loading Vendors List')
        this.vendors = Vendors.query()
      },
      getEquipment: function() {
        console.log('Loading Equipment List')
        this.equipment = Equipment.query()
      },
      getEquipmentDetails: function(id) {
        this.equipmentDetail = null
        angular.forEach(this.equipment, function(v,k){
          if (v.id === id) {
            this.equipmentDetail = v
          }
        },this)
      },
      myCibo: function() {
        alert('myCibo')
      }

    });
    
  });

})();
