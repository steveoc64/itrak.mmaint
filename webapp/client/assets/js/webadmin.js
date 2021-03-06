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


  // Admin version of Spares List
  angular.module('itrak')
    .controller('adminSparesCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,    
      spares,sites,vendors
    ){     

    console.log('Running Admin Spares controller',$stateParams)

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      spares: spares,
      sites: sites,
      vendors: vendors,

      getSiteMapURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },

      changed: function(spare) {
        console.log('Changed',spare.id)
        spare.$save({id: spare.id})
      }

    }) // extend this 

  })

  // Admin version of Consumables List
  angular.module('itrak')
    .controller('adminConsumablesCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,    
      consumables,sites,vendors
    ){     

    console.log('Running Admin Consumables controller',$stateParams)

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      consumables: consumables,
      sites: sites,
      vendors: vendors,

      getSiteMapURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },

      changed: function(part) {
        console.log('Changed',part.id)
        part.$save({id: part.id})
      }

    }) // extend this 

  })

  // Admin version of EquipTypes List
  angular.module('itrak')
    .controller('adminEquipTypesCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,    
      equiptypes
    ){     

    console.log('Running Admin EquipTypes controller',$stateParams)

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      equiptypes: equiptypes,

      changed: function(et) {
        console.log('Changed',et)
        et.$save({id: et.id})
      }

    }) // extend this 

  })

  // Admin version of Vendor List
  angular.module('itrak')
    .controller('adminVendorCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,    
      vendors
    ){     

    console.log('Running Admin EquipTypes controller',$stateParams)

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      vendors: vendors,
      ratings: [
        {id: 1, name: 'A+'},
        {id: 2, name: 'Good'},
        {id: 3, name: 'Average'},
        {id: 4, name: 'Poor'},
        {id: 5, name: 'Avoid'},
      ],

      changed: function(vendor) {
        console.log('Changed',vendor)
        vendor.$save({id: vendor.id})
      }

    }) // extend this 

  })

  // Admin version of Equipment Detail
  angular.module('itrak')
    .controller('adminEquipmentDetCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,
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
        this.equipment.$save()
        FoundationApi.publish('equipment','reload')
      },
      getSiteMapURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },
    });  // Extend this  
  });

  // Admin version of Equipment Detail
  angular.module('itrak')
    .controller('adminPeopleCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,
      people,sites,roles
    ){     

    console.log('Running Admin People controller',$stateParams)
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      people: people,
      roles: roles,
      sites: sites,

      changed: function(person) {
        console.log('people ?',person.id)
        person.$save({id: person.id})
        FoundationApi.publish('people','reload')
      },
    });  // Extend this  
  });

  // Admin version of Site List
  angular.module('itrak')
    .controller('adminSiteCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,
      sites
    ){     

    console.log('Running Admin Site controller',$stateParams)
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      sites: sites,

      getSiteMapURI: function(m) {
        return "https://www.google.com/maps?q="+encodeURIComponent(m)
      },

      changed: function(site) {
        console.log('site changed',site.id)
        site.$save({id: site.id})
        FoundationApi.publish('sites','reload')
      },
    });  // Extend this  
  });

  // Admin version of Task List
  angular.module('itrak')
    .controller('adminTaskCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,
      tasks
    ){     

    console.log('Running Admin Task controller',$stateParams)
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      tasks: tasks,
      showInstructions: false,

      changed: function(task) {
        console.log('Task changed',site.id)
        task.$save({id: task.id})
        FoundationApi.publish('tasks','reload')
      },
    });  // Extend this  
  });

  // Admin version of Site Task List
  angular.module('itrak')
    .controller('adminSiteTaskCtrl', function(
      $state, $stateParams,
      FoundationApi,
      loginState,
      tasks, site
    ){     

    console.log('Running Admin Site Task controller',$stateParams)
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      tasks: tasks,
      site: site,

      changed: function(task) {
        console.log('Something changed ...')
      },
    });  // Extend this  
  });

})();
