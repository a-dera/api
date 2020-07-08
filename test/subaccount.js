var DuniaPay = require('../index')(process.env.KEY)
  , mocha = require('mocha')
  , expect = require('chai').expect
  ;


describe("DuniaPay Subaccount", function() {

  var subaccount_id, subaccount_code;

  // New Subaccount
  it("should create a new subaccount", function(done) {
    DuniaPay.subaccount.create({
      business_name: 'Super Cool Inc',
      settlement_bank: 'Access Bank',
      account_number: '0193274682',
      percentage_charge: 18.2
    })
    .then(function(body){
      expect(body).to.have.property('data');
      expect(body.data).to.have.property('id');
      expect(body.data).to.have.property('subaccount_code');

      subaccount_id = body.data.id;
      subaccount_code = body.data.subaccount_code;
      done();
    })
    .catch(function(error){
      return done(error);
    });
  });

  // Update Subaccount
  it("should update a subaccount", function(done) {
    DuniaPay.subaccount.update(subaccount_code, {
      primary_contact_email: 'wale@obo.com',
      percentage_charge: 98.9
    })
    .then(function(body){
      expect(body).to.have.property('data');
      expect(body.data).to.have.property('id');
      done();
    })
    .catch(function(error){
      return done(error);
    });
  });

  // Fetch Subaccount
  it("should get details of a subaccount", function(done) {
    DuniaPay.subaccount.get(subaccount_code)
    .then(function(body){
      expect(body).to.have.property('data');
      expect(body.data).to.have.property('id');
      done();
    })
    .catch(function(error){
      return done(error);
    });
  });

  // list Banks
  it("should list supported banks", function(done) {
    DuniaPay.subaccount.listBanks()
    .then(function(body){
      expect(body).to.have.property('data');
      expect(body.data).to.be.instanceof(Array);
      done();
    })
    .catch(function(error){
      return done(error);
    });
  });


  // List Subaccounts
  it("should list subaccounts", function(done) {
    DuniaPay.subaccount.list()
    .then(function(body){
      expect(body).to.have.property('data');
      expect(body.data).to.be.instanceof(Array);
      done();
    })
    .catch(function(error){
      return done(error);
    });
  });

});
