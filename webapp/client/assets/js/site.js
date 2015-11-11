;(function(){
  // siteCtrl
  'use strict';

  console.log('Loading Site Controller')

  // Remote resource for Site transactions
  angular.module('itrak').factory('SiteServer', function($resource, ServerName){
    return $resource(ServerName+'/site/:id');
  })

  // Controller for the Site Page
  angular.module('itrak').controller('siteCtrl', function($state, loginState, SiteServer){     
    console.log('Running Site controller')
    if (!loginState.loggedIn) {
      $state.go('login')
    }

    angular.extend(this, {
      loginState: loginState,
      siteList: [],
    })

    // Get the site list
    this.siteList = SiteServer.query()
  });

})();
