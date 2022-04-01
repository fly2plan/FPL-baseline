const avro = require('avsc');

const orgEventType = avro.Type.forSchema({
  type: 'record',
  fields: [
    {
      name: 'id',
      type: 'string'
    },
    {
      name: 'name',
      type: 'string'
    },
    {
      name: 'pk',
      type: 'string'
    }
  ]
});

const workgroupEventType = avro.Type.forSchema({
  type: 'record',
  fields: [
    {
      name: 'id',
      type: 'string'
    },
    {
      name: 'shieldContractAddress',
      type: 'string'
    },
    {
      name: 'members',
      type: {
        type: "array",
        items: {
          type: 'record',
          name: 'member',
          fields: [
            {
              name: 'id',
              type: 'string'
            },
            {
              name: 'pk',
              type: 'string'
            }
          ]
        },
        default: []
      }
    },
    {
      name: 'workflows',
      type: {
        type: 'array',
        items: 'string'
      },
      default: []
    }
  ]
});

const workflowSyncFPLEventType = avro.Type.forSchema({
  type: 'record',
  fields: [
    {
      name: 'workgroupId',
      type: 'string'
    },
    {
      name: 'workflowId',
      type: 'string'
    },
    {
      name: 'workstepId',
      type: 'string'
    },
    {
      name: 'SYNC',
      type: {
        type: 'record',
        name: 'Sync',
        fields: [
          {
            name: 'fpl',
            type: {
              type: 'record',
              name: 'FPL',
              fields: [
                {
                  name: 'raw',
                  type: {
                    type: 'record',
                    name: 'Raw',
                    fields: [
                      {
                        name: 'ai',
                        type: 'string'
                      },
                      {
                        name: 'fr',
                        type: 'string'
                      },
                      {
                        name: 'tof',
                        type: 'string'
                      },
                      {
                        name: 'noa',
                        type: 'string'
                      },
                      {
                        name: 'toa',
                        type: 'string'
                      },
                      {
                        name: 'wtc',
                        type: 'string'
                      },
                      {
                        name: 'eandc',
                        type: {
                          type: 'record',
                          name: 'EquipmentAndCapabilities',
                          fields: [
                            {
                              name: 'rc_n_aae',
                              type: {
                                type: 'record',
                                name: 'RC_N_AAE',
                                fields: [
                                  {
                                    name: 's',
                                    type: 'string'
                                  },
                                  {
                                    name: 'a',
                                    type: 'string'
                                  }
                                ]
                              }
                            },
                            {
                              name: 'surveillance',
                              type: {
                                type: 'record',
                                name: 'Surveillance',
                                fields: [
                                  {
                                    name: 's',
                                    type: 'string'
                                  },
                                  {
                                    name: 'ssr',
                                    type: {
                                      type: 'record',
                                      name: 'SSR',
                                      fields: [
                                        {
                                          name: 'mode_a_c',
                                          type: {
                                            type: 'record',
                                            name: 'Mode_a_c',
                                            fields: [
                                              {
                                                name: 'a',
                                                type: 'string'
                                              }
                                            ]
                                          }
                                        },
                                        {
                                          name: 'mode_s',
                                          type: {
                                            type: 'record',
                                            name: 'Mode_s',
                                            fields: [
                                              {
                                                name: 'a',
                                                type: 'string'
                                              }
                                            ]
                                          }
                                        }
                                      ]
                                    }
                                  },
                                  {
                                    name: 'ads',
                                    type: {
                                      type: 'record',
                                      name: 'ADS',
                                      fields: [
                                        {
                                          name: 'c',
                                          type: {
                                            type: 'record',
                                            name: 'C',
                                            fields: [
                                              {
                                                name: 'a',
                                                type: 'string'
                                              }
                                            ]
                                          }
                                        },
                                        {
                                          name: 'b',
                                          type: {
                                            type: 'record',
                                            name: 'B',
                                            fields: [
                                              {
                                                name: 'a',
                                                type: 'string'
                                              }
                                            ]
                                          }
                                        }
                                      ]
                                    }
                                  }
                                ]
                              }
                            }
                          ]
                        }
                      },
                      {
                        name: 't',
                        type: 'string'
                      },
                      {
                        name: 'da',
                        type: 'string'
                      },
                      {
                        name: 'cs',
                        type: 'string'
                      },
                      {
                        name: 'cl',
                        type: 'string'
                      },
                      {
                        name: 'r',
                        type: 'string'
                      },
                      {
                        name: 'de',
                        type: 'string'
                      },
                      {
                        name: 'teet',
                        type: 'string'
                      },
                      {
                        name: 'ada',
                        type: 'string'
                      },
                      {
                        name: 'adas',
                        type: {
                          type: 'array',
                          items: 'string'
                        }
                      },
                      {
                        name: 'oi',
                        type: 'string'
                      },
                      {
                        name: 'e',
                        type: 'string'
                      },
                      {
                        name: 'pob',
                        type: 'string'
                      },
                      {
                        name: 'er',
                        type: 'string'
                      },
                      {
                        name: 'se',
                        type: 'string'
                      },
                      {
                        name: 'j',
                        type: 'string'
                      },
                      {
                        name: 'd',
                        type: 'string'
                      },
                      {
                        name: 'acm',
                        type: 'string'
                      },
                      {
                        name: 're',
                        type: 'string'
                      },
                      {
                        name: 'pn',
                        type: 'string'
                      }
                    ]
                  },
                },
                {
                  name: 'proof',
                  type: {
                    type: 'record',
                    name: 'Proof',
                    fields: [
                      {
                        name: 'a',
                        type: {
                          type: 'array',
                          items: 'string'
                        }
                      },
                      {
                        name: 'b',
                        type: {
                          type: 'array',
                          items: {
                            type: 'array',
                            name: 'b_item',
                            items: 'string'
                          }
                        }
                      },
                      {
                        name: 'c',
                        type: {
                          type: 'array',
                          items: 'string'
                        }
                      }
                    ]
                  },
                },
                {
                  name: 'inputs',
                  type: {
                    type: 'record',
                    name: 'Inputs',
                    fields: [
                      {
                        name: 'hash',
                        type: 'string'
                      },
                      {
                        name: 'sig',
                        type: 'string'
                      },
                      {
                        name: 'pk',
                        type: 'string'
                      }
                    ]
                  }
                }
              ]
            },
          },
        ]
      },
    }
  ]
});

const workflowSyncACKEventType = avro.Type.forSchema({
  type: 'record',
  fields: [
    {
      name: 'workgroupId',
      type: 'string'
    },
    {
      name: 'workflowId',
      type: 'string'
    },
    {
      name: 'workstepId',
      type: 'string'
    },
    {
      name: 'SYNC',
      type: {
        type: 'record',
        name: 'Sync',
        fields: [
          {
            name: 'ack',
            type: {
              type: 'record',
              name: 'ACK',
              fields: [
                {
                   name: 'raw',
                   type: 'string'
                },
                {
                  name: 'proof',
                  type: {
                    type: 'record',
                    name: 'Proof',
                    fields: [
                      {
                        name: 'a',
                        type: {
                          type: 'array',
                          items: 'string'
                        }
                      },
                      {
                        name: 'b',
                        type: {
                          type: 'array',
                          items: {
                            type: 'array',
                            name: 'b_item',
                            items: 'string'
                          }
                        }
                      },
                      {
                        name: 'c',
                        type: {
                          type: 'array',
                          items: 'string'
                        }
                      }
                    ]
                  },
                },
                {
                  name: 'inputs',
                  type: {
                    type: 'record',
                    name: 'Inputs',
                    fields: [
                      {
                        name: 'fpl_hash',
                        type: 'string'
                      },
                      {
                        name: 'ao_sig',
                        type: 'string'
                      },
                      {
                        name: 'ao_pk',
                        type: 'string'
                      },
                      {
                        name: 'ack_hash',
                        type: 'string'
                      },
                      {
                        name: 'nm_sig',
                        type: 'string'
                      },
                      {
                        name: 'nm_pk',
                        type: 'string'
                      }
                    ]
                  }
                }
              ]
            },
          },
        ]
      },
    }
  ]
});

module.exports = {
  orgEventType,
  workgroupEventType,
  workflowSyncFPLEventType,
  workflowSyncACKEventType
}