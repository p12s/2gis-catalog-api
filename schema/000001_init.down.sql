UPDATE schema_migrations SET dirty = false;

DROP TABLE IF EXISTS rubric CASCADE;

DROP TABLE IF EXISTS city CASCADE;
DROP TABLE IF EXISTS street CASCADE;

DROP TABLE IF EXISTS building CASCADE;

DROP TABLE IF EXISTS company CASCADE;
DROP TABLE IF EXISTS company_phone CASCADE;
DROP TABLE IF EXISTS company_rubric CASCADE;
DROP TABLE IF EXISTS phone CASCADE;
