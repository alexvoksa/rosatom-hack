export default (reqObj) => `
query a {
  products {
    guid
    okei_id
    price
    unit {
      id
    }
  }
}
`