import Mixin from '@ember/object/mixin';
import { inject as service } from '@ember/service';
import { scheduleOnce } from '@ember/runloop';

export default Mixin.create({
  routeClass: service(),

  activate() {
    this._super(...arguments);
    scheduleOnce('afterRender', this, this._updateRouteClass);
  },

  _updateRouteClass() {
    this.routeClass.routeName = this.routeName;
  }
});