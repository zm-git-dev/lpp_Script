#!/home/nfs/SOFTWARE/bin/nextflow
params.query = "$HOME/sample.fa"
params.db = "$HOME/tools/blast-db/pdb/pdb"
params.out = "./result.txt"
params.chunkSize = 1000
db_name = file(params.db)

 
/*
 * Given the query parameter creates a channel emitting the query fasta file(s),
 * the file is split in chunks containing as many sequences as defined by the parameter 'chunkSize'.
 * Finally assign the result channel to the variable 'fasta'
 */
Channel
    .fromPath(params.query)
    .splitFasta(by: params.chunkSize)
    .set { fasta }
 
/*
 * Executes a BLAST job for each chunk emitted by the 'fasta' channel
 * and creates as output a channel named 'top_hits' emitting the resulting
 * BLAST matches 
 */
process index {
    executor 'pbs'
    cpus 32  
    clusterOptions  "   -d $PWD "
    input:
                file db_name

    
    output:
                file "ref.mmi" into MinimapIndex

    script:
	"""
		minimap2 -d ref.mmi $db_name
	"""
}
process Minimap {
    executor 'pbs'
    cpus 1
    clusterOptions  " -d $PWD   -l nodes=1:ppn=1 -V"
    input:
    	file Minimapgenome  from MinimapIndex.first()
    	file 'query.fa' from fasta
 
    output:
    file 'result.txt' into top_hits
 
  

	"""	
	minimap2 -x splice $Minimapgenome query.fa 2>&1 |grep -vF '['| cut -f 1,2,3,4,5,6,7,8,9,11|  Minimap_First.py >result.txt

    """
}
 
 
/*
 * Each time a file emitted by the 'top_hits' channel an extract job is executed
 * producing a file containing the matching sequences
 */

/*
 * Collects all the sequences files into a single file
 * and prints the resulting file content when complete
 */
top_hits.collectFile(name: params.out)

