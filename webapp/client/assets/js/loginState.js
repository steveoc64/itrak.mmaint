(function(){
	// loginState
	'use strict';

	console.log("Loading LoginState")

	angular.module('itrak').service('loginState', function($http){
		angular.extend(this, {
			loggedIn: false,
			toggle: function() {
				this.loggedIn = !this.loggedIn
				console.log("Login state toggled to", this.loggedIn)
			},
			login: function(username,passwd) { 
				console.log("logging in with credentials",
					username,
					passwd)

				$http.put('http://localhost:8082/login', {u: username, p: passwd})
			},
			logout: function() {
				this.loggedIn = false
			},
		})

		console.log('loginState init to',this)
	})

})();