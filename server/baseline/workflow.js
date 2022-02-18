const express = require('express')
const router = express.Router()

const { v4: uuidv4 } = require('uuid');

const { workflowSyncRouter } = require('./workflowSync');

const { workflowExists } = require('./workflowRegistry');
const { workgroupExists } = require('./workgroupRegistry');

router.use((req, res, next) => {
    if (req.body.workgroupId !== undefined) {
        if (!workgroupExists(req.body.workgroupId)) {
            return res.sendStatus(403);
        };
    } else {
        return res.sendStatus(403);
    };
    if (req.body.workflowId !== undefined) {
        if (!workflowExists(req.body.workflowId)) {
            return res.sendStatus(403);
        }
    } else {
        let id = uuidv4();
        res.locals.id = id;
    }
    next();
})


router.use('/sync', workflowSyncRouter);

module.exports = {
    workflowRouter: router,
}