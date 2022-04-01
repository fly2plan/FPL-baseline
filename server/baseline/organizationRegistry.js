const orgRegistry = new Map()

const organizationExists = (id) => {
  return orgRegistry.has(id)
}

const insertOrg = (organization) => {
  if (orgRegistry.has(organization.id)) {
    console.log(`organization with id ${organization.id} already exists`)
    return;
  }

  console.log('adding new organization ', organization)
  orgRegistry.set(organization.id, {
    id: organization.id,
    name: organization.name,
    pk: organization.pk
  })
}

const getOrg = (id) => {
  if (organizationExists(id)) {
    return orgRegistry.get(id)
  }
}

module.exports = {
  orgRegistry,
  organizationExists,
  insertOrg,
  getOrg,
}