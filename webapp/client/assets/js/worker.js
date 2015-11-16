;(function(){
  // workerCtrl
  'use strict';

  console.log('Loading Worker Controller')

  // Controller for the Admin Page
  angular.module('itrak')
    .controller('workerCtrl', function($state, $stateParams, loginState){     
    
    console.log('Running Worker controller')

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


  // Controller for the Worker EStop
  angular.module('itrak').controller('workerStopCtrl', 
    function($state, loginState, siteEquip){     

    console.log('Running Worker Stop controller')

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    console.log(siteEquip)
    angular.extend(this, {
      loginState: loginState,
      siteEquip: siteEquip,
    })

  });
  // Controller for the Worker EStop
  angular.module('itrak').controller('workerPStopCtrl', 
    function($state, loginState, siteEquip){     

    console.log('Running Worker Preventative Maint controller')

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    console.log(siteEquip)
    angular.extend(this, {
      loginState: loginState,
      siteEquip: siteEquip,
    })

  });

  // Controller for the Worker Equip
  angular.module('itrak').controller('workerEquipmentCtrl', 
    function($state, loginState, siteEquip){     

    console.log('Running Worker Equip controller')

    if (!loginState.loggedIn) {
      $state.go('login')
    }

    console.log(siteEquip)
    angular.extend(this, {
      loginState: loginState,
      siteEquip: siteEquip,
    })

  });



})();
