import Route from '@ember/routing/route';
import { service } from '@ember/service';

export default class ApplicationRoute extends Route {
  // Your code goes here
  @service router;
    beforeModel() {
        console.log('beforeModel');
        //check if token is present in session or local storage and redirect to dashboard
        if (sessionStorage.getItem('token')) {
        //redirect to dashboard
            this.router.transitionTo('dashboard');
            return;
        }
        this.router.transitionTo('login');
    }
}