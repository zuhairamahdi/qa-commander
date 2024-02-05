import EmberObject from '@ember/object';
import RouteClassMixin from 'frontend/mixins/route-class';
import { module, test } from 'qunit';

module('Unit | Mixin | route-class', function () {
  // TODO: Replace this with your real tests.
  test('it works', function (assert) {
    let RouteClassObject = EmberObject.extend(RouteClassMixin);
    let subject = RouteClassObject.create();
    assert.ok(subject);
  });
});
