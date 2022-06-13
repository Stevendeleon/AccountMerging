const accounts = require('./accounts.json');
const mAccounts = require('./index.js');

const mockData = [
    {
        applications: [ '1', '2', '3' ],
        emails: [
            'a@gmail.com',
            'b@gmail.com',
            'a@yahoo.com',
            'a@gmail.com',
            'a@yahoo.com'
        ],
        name: 'A'
    },
    {
        applications: [ '1' ],
        emails: [ 'c@gmail.com', 'd@gmail.com' ],
        name: 'C'
    }
];


describe('mergeAccounts', () => {
    it('should return an array matching the mock data', () => {
        expect(mAccounts.mergeAccounts(accounts)).toEqual(expect.any(Array));
        expect(mAccounts.mergeAccounts(accounts)).toEqual(mockData);
    })
})

