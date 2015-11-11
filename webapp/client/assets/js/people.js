;(function(){
  // peopleCtrl
  'use strict';

  console.log('Loading People Controller')

  // Remote resource for People transactions
  angular.module('itrak').factory('PeopleServer', function($resource, ServerName){
    return $resource(ServerName+'/people/:id');
  })

  // Controller for the SiteMgr Page
  angular.module('itrak').controller('peopleCtrl', function($state, loginState, PeopleServer){     
    console.log('Running People controller')

    if (!loginState.loggedIn) {
      $state.go('login')
    } 

    angular.extend(this, {
      loginState: loginState,
      peopleList: [],
    })

    // Now fetch the initial people list
    this.peopleList = PeopleServer.query()

  });

})();
