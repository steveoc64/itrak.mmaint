;(function(){
  // adminCtrl
  'use strict';

  console.log('Loading Admin Controller')

  // Controller for the Admin Page
  angular.module('itrak')
    .controller('adminCtrl', function(
      $state, $stateParams, loginState   
    ){     
    console.log('Running Admin controller')

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
      menuSelection: '',
      params: $stateParams,
      allSites: false,
    });
    
  });

  // Admin version of Equipment List
  angular.module('itrak')
    .controller('adminEquipmentCtrl', function(
      $state, $stateParams,
      FoundationApi,Equipment,
      loginState,    
      equipments,sites,vendors
    ){     

    console.log('Running Admin Equipment controller',$stateParams)

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      equipmentTypes: [
        {id: 1, name: 'Consumables'},
        {id: 2, name: 'Mech'},
        {id: 3, name: 'Spare Parts'},
      ],

      allSites: false,
      selectedSites: [],
      siteEquips: [],
      mainSiteEquips: [],      
      equipments: equipments,
      sites: sites,
      vendors: vendors,

      siteSelected: function(id) {
        console.log('Toggled site',id)
        this.rebuildFilteredEquip()
      },

      toggleAllSites: function() {
        console.log('ToggleAll')
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
        angular.forEach(this.equipments, function(v,k){
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

      getSiteMapURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },

      changed: function(v) {
        console.log('Changed',v)
      }

    }) // extend this 

    var vm = this
    FoundationApi.subscribe('equipment', function(v) {
      console.log('MSG: equipment',v)
      switch(v) {
        case 'reload':
          Equipment.query().$promise.then(function(result){
            vm.equipments = result
            vm.rebuildFilteredEquip()
          })
          break
      }
    })

  })

  // Admin version of Equipment Detail
  angular.module('itrak')
    .controller('adminEquipmentDetCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState, $scope,
      equipment,sites,vendors,subparts
    ){     

    console.log('Running Admin Equipment Det controller',$stateParams)
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      equipmentTypes: [
        {id: 1, name: 'Consumables'},
        {id: 2, name: 'Mech'},
        {id: 3, name: 'Spare Parts'},
      ],
      allSites: false,
      selectedSites: [],
      siteEquips: [],
      mainSiteEquips: [],
      equipment: equipment,
      sites: sites,
      vendors: vendors,
      subparts: subparts,

      changed: function() {
        equipment.$save()
        FoundationApi.publish('equipment','reload')
      },
      getSiteMapURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },
    });  // Extend this  
  });

})();
