var DuniaPay = require('../index')(process.env.KEY)
  , mocha = require('mocha')
  , expect = require('chai').expect
  ;

describe("DuniaPay Settlements", function() {

  // Fetch Settlements
  it("should fetch settlements", function(done) {
    DuniaPay.settlements.fetch()
      .then(function(body){
        expect(body).to.have.property('status');
        expect(body).to.have.property('message');
        
        done();
      }).catch(function(error){
        return done(error);
      });
  });
});
