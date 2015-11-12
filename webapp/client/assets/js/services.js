;(function(){
  // Define all services in here
  'use strict';

  console.log('Loading Services')

  // People
  angular.module('itrak').factory('PeopleServer', function($resource, ServerName){
    return $resource(ServerName+'/people/:id');
  })


})();
