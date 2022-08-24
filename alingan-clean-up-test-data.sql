-- Run to clean up test data

delete from core."owner";
delete from core.agent;
delete from core.store;
delete from core.product;
delete from core.product_image;
delete from core."transaction";
delete from core.transaction_item;
delete from core.transaction_receipt;