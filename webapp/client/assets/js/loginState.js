(function(){
	// loginState
	'use strict';

	console.log("Loading LoginState")

	angular.module('itrak').service('loginState', function(){
		angular.extend(this, {
			loggedIn: false,
			toggle: function() {
				this.loggedIn = !this.loggedIn
				console.log("Login state toggled to", this.loggedIn)
			},
			login: function() { 
				this.loggedIn = true
			},
			logout: function() {
				this.loggedIn = false
			}
		})

		console.log('loginState init to',this)
	})

})();