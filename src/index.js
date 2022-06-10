const accounts = require('../accounts.json');

function mergeAccounts(accounts) {

  let mergedAccounts = {};
  let mergedAccountsList = [];

  for (const account of accounts) {
    if (mergedAccounts[account.name]) {
      mergedAccounts[account.name].applications.push(account.application);
      mergedAccounts[account.name].emails = mergedAccounts[account.name].emails.concat(account.emails);
    } else {
      mergedAccounts[account.name] = {
        applications: [account.application],
        emails: account.emails,
        name: account.name
      }
    }
  }

  for (const key in mergedAccounts) {
    mergedAccountsList.push(mergedAccounts[key]);
  }

  return mergedAccountsList;
}

const accountsList = mergeAccounts(accounts);
console.log(accountsList);
